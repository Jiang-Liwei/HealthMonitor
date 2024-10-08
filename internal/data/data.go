package data

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/extra/redisotel/v9"

	"healthmonitor/ent"
	"healthmonitor/internal/conf"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewBloodStatusRepo,
	NewAdminUserRepo,
)

// Data holds the database and Redis clients.
type Data struct {
	db  *ent.Client
	rdb *redis.Client
}

// NewData initializes the Data structure with Redis and Ent clients.
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	logHelper := log.NewHelper(logger)

	// Initialize Ent client with tracing and logging
	client, err := NewEntClient(c, logHelper)
	if err != nil {
		return nil, nil, err
	}

	// Initialize Redis client with tracing
	rdb := NewRedis(c)

	d := &Data{
		db:  client,
		rdb: rdb,
	}

	cleanup := func() {
		logHelper.Info("closing the data resources")
		if err := d.db.Close(); err != nil {
			logHelper.Error(err)
		}
		if err := d.rdb.Close(); err != nil {
			logHelper.Error(err)
		}
	}

	return d, cleanup, nil
}

// NewRedis initializes the Redis client with tracing support.
func NewRedis(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})

	// 使用 redisotel 进行 tracing 集成
	if err := redisotel.InstrumentTracing(rdb); err != nil {
		log.Fatalf("failed to instrument redis client: %v", err)
	}

	return rdb
}

// NewEntClient initializes the Ent client with tracing and logging.
func NewEntClient(c *conf.Data, logHelper *log.Helper) (*ent.Client, error) {
	drv, err := sql.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	if err != nil {
		logHelper.Errorf("failed opening connection to database: %v", err)
		return nil, err
	}

	// Add tracing and logging to the SQL driver
	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		logHelper.WithContext(ctx).Info(i...)
		tracer := otel.Tracer("ent.")
		kind := trace.SpanKindServer
		_, span := tracer.Start(ctx,
			"Query",
			trace.WithAttributes(
				attribute.String("sql", fmt.Sprint(i...)),
			),
			trace.WithSpanKind(kind),
		)
		span.End()
	})

	client := ent.NewClient(ent.Driver(sqlDrv))

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		logHelper.Errorf("failed creating schema resources: %v", err)
		return nil, err
	}

	return client, nil
}
