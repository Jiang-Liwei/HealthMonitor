package admin

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	v1 "healthmonitor/api/admin/v1"
	"healthmonitor/ent/bloodstatusrecord"
	"healthmonitor/internal/biz/api"
	"sync"
	"time"
)

// DashboardRepo 后台首页
type DashboardRepo interface {
}

type DashboardUsecase struct {
	repo api.BloodStatusRepo
	log  *log.Helper
}

func NewDashboardUsecase(repo api.BloodStatusRepo, logger log.Logger) *DashboardUsecase {
	return &DashboardUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DashboardUsecase) Dashboard(ctx context.Context, userID uuid.UUID) (*v1.DashboardReply, error) {
	var (
		allNum, goodNum, badNum                                                    int
		last7DaysAllNum, last7DaysGoodNum, last7DaysBadNum                         int
		previousLast7DaysAllNum, previousLast7DaysGoodNum, previousLast7DaysBadNum int
	)

	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// 计算最近7天和前一周期的开始和结束时间
	last7DaysStartTime := todayStart.Add(-6 * 24 * time.Hour).UnixMilli()
	last7DaysEndTime := todayStart.Add(24*time.Hour - time.Millisecond).UnixMilli()
	previousLast7DaysStartTime := todayStart.Add(-13 * 24 * time.Hour).UnixMilli()
	previousLast7DaysEndTime := todayStart.Add(-7*24*time.Hour + 24*time.Hour - time.Millisecond).UnixMilli()

	var wg sync.WaitGroup
	var errList []error
	var mu sync.Mutex

	performQuery := func(name string, result *int, fn func() (int, error)) {
		defer wg.Done()
		num, err := fn()
		if err != nil {
			mu.Lock()
			errList = append(errList, fmt.Errorf("%s query failed: %w", name, err))
			mu.Unlock()
		}
		*result = num
	}

	wg.Add(9)
	go performQuery("allNum", &allNum, func() (int, error) { return uc.repo.CountByUserID(ctx, userID) })
	go performQuery("goodNum", &goodNum, func() (int, error) { return uc.repo.CountByStatusSummaryAndUserID(ctx, userID, "perfect") })
	go performQuery("badNum", &badNum, func() (int, error) { return uc.repo.CountByStatusSummaryAndUserID(ctx, userID, "very bad") })
	go performQuery("last7DaysAllNum", &last7DaysAllNum, func() (int, error) { return uc.repo.GetPeriodNum(ctx, userID, last7DaysStartTime, last7DaysEndTime) })
	go performQuery("last7DaysGoodNum", &last7DaysGoodNum, func() (int, error) {
		return uc.repo.GetByStatusSummaryAndPeriodNum(ctx, userID, last7DaysStartTime, last7DaysEndTime, "perfect")
	})
	go performQuery("last7DaysBadNum", &last7DaysBadNum, func() (int, error) {
		return uc.repo.GetByStatusSummaryAndPeriodNum(ctx, userID, last7DaysStartTime, last7DaysEndTime, "very bad")
	})
	go performQuery("previousLast7DaysAllNum", &previousLast7DaysAllNum, func() (int, error) {
		return uc.repo.GetPeriodNum(ctx, userID, previousLast7DaysStartTime, previousLast7DaysEndTime)
	})
	go performQuery("previousLast7DaysGoodNum", &previousLast7DaysGoodNum, func() (int, error) {
		return uc.repo.GetByStatusSummaryAndPeriodNum(ctx, userID, previousLast7DaysStartTime, previousLast7DaysEndTime, "perfect")
	})
	go performQuery("previousLast7DaysBadNum", &previousLast7DaysBadNum, func() (int, error) {
		return uc.repo.GetByStatusSummaryAndPeriodNum(ctx, userID, previousLast7DaysStartTime, previousLast7DaysEndTime, "very bad")
	})

	wg.Wait()

	if len(errList) > 0 {
		for _, err := range errList {
			uc.log.Error(err)
		}
	}

	// 计算三种数据的增长率
	computeChangeRatio := func(current, previous int) float32 {
		if previous > 0 {
			return float32(current-previous) / float32(previous) * 100
		}
		return 0
	}

	allChangeRatio := computeChangeRatio(last7DaysAllNum, previousLast7DaysAllNum)
	goodChangeRatio := computeChangeRatio(last7DaysGoodNum, previousLast7DaysGoodNum)
	badChangeRatio := computeChangeRatio(last7DaysBadNum, previousLast7DaysBadNum)

	// 最近6个月的数据
	buildMonthDetails := func(ctx context.Context, userID uuid.UUID, status bloodstatusrecord.StatusSummary) []*v1.MonthDetails {
		var monthDetails []*v1.MonthDetails
		for i := 0; i < 6; i++ {
			monthEnd := time.Date(now.Year(), now.Month()-time.Month(i), 1, 0, 0, 0, 0, now.Location()).AddDate(0, 1, -1).Add(23*time.Hour + 59*time.Minute + 59*time.Second)
			monthStart := time.Date(monthEnd.Year(), monthEnd.Month(), 1, 0, 0, 0, 0, monthEnd.Location())

			var monthData int
			var err error
			if status != "" {
				monthData, err = uc.repo.GetByStatusSummaryAndPeriodNum(ctx, userID, monthStart.UnixMilli(), monthEnd.UnixMilli(), status)
			} else {
				monthData, err = uc.repo.GetPeriodNum(ctx, userID, monthStart.UnixMilli(), monthEnd.UnixMilli())
			}

			if err != nil {
				uc.log.Error(fmt.Errorf("failed to get count for %s in month %d: %w", status, i+1, err))
			}
			monthDetails = append(monthDetails, &v1.MonthDetails{
				Num:   int64(monthData),
				Month: monthStart.Format("2006-01"),
			})
		}
		return monthDetails
	}

	// all, good, bad 三种 CardList
	cardList := []*v1.CardList{
		{
			Num:          int64(last7DaysAllNum),
			ChangeRatio:  allChangeRatio,
			Category:     "all",
			MonthDetails: buildMonthDetails(ctx, userID, ""),
		},
		{
			Num:          int64(last7DaysGoodNum),
			ChangeRatio:  goodChangeRatio,
			Category:     "good",
			MonthDetails: buildMonthDetails(ctx, userID, "perfect"),
		},
		{
			Num:          int64(last7DaysBadNum),
			ChangeRatio:  badChangeRatio,
			Category:     "bad",
			MonthDetails: buildMonthDetails(ctx, userID, "very bad"),
		},
	}

	// 获取最近7条记录
	recentRecords, err := uc.repo.GetRecentRecords(ctx, userID, 7)
	if err != nil {
		uc.log.Error(fmt.Errorf("failed to get recent records: %w", err))
		return nil, err
	}

	var lastSevenRecords []*v1.LastSevenRecords
	for _, record := range recentRecords {
		lastSevenRecords = append(lastSevenRecords, &v1.LastSevenRecords{
			RecordDate:        int64(record.RecordDate),
			StatusSummary:     string(record.StatusSummary),
			SystolicPressure:  uint64(record.SystolicPressure),
			DiastolicPressure: uint64(record.DiastolicPressure),
			Pulse:             uint64(record.Pulse),
			Mood:              string(record.Mood),
			Id:                record.ID.String(),
		})
	}

	return &v1.DashboardReply{
		CardList:         cardList,
		LastSevenRecords: lastSevenRecords,
	}, nil
}
