package actions

import (
	"context"
	"log"

	"ai-gallery/ent/migrate"
	"ai-gallery/service"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use: "migrate",
		Run: MigrateDB,
	}
	rootCmd.AddCommand(cmd)
}

func MigrateDB(cmd *cobra.Command, args []string) {
	env, err := cmd.Flags().GetString("env")
	if err != nil {
		log.Fatalf("get env failed: %v\n", err)
	}

	c := service.LoadConfig(env)
	client, err := service.InitDB(c)

	if err := client.Debug().Schema.Create(
		context.Background(),
		schema.WithHooks(func(next schema.Creator) schema.Creator {
			return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
				return next.Create(ctx, tables...)
			})
		}),
		migrate.WithForeignKeys(false),
	); err != nil {
		log.Fatal(err.Error())
	}
	log.Println("migrate successfully")
}
