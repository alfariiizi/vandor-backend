package seeder

//
// import (
// 	"context"
//
// 	"github.com/alfariiizi/go-service/ent/db"
// 	"go.uber.org/fx"
// )
//
// type Seeder struct {
// 	db *db.Client
// }
//
// func NewSeeder(lc fx.Lifecycle, db *db.Client) *Seeder {
// 	seeder := &Seeder{
// 		db: db,
// 	}
//
// 	lc.Append(fx.Hook{
// 		OnStart: func(ctx context.Context) error {
// 			return nil
// 		},
// 	})
//
// 	return seeder
// }
//
// func GenerateAdmin(db *db.Client) error {
// 	// cfg := config.GetConfig()
//
// 	tx, err := db.Tx(context.Background())
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			_ = tx.Rollback()
// 			panic(r)
// 		}
// 	}()
//
// 	// product, err := tx.Product.Create().
// 	// 	SetName("Pentaverses Auth").
// 	// 	SetSlug("pentaverses_auth").
// 	// 	Save(context.Background())
// 	// if err != nil {
// 	// 	_ = tx.Rollback()
// 	// 	panic(err)
// 	// }
// 	//
// 	// role, err := tx.Role.Create().
// 	// 	SetName("Super Admin").
// 	// 	SetSlug("super_admin").
// 	// 	SetProductID(product.ID).
// 	// 	Save(context.Background())
// 	// if err != nil {
// 	// 	_ = tx.Rollback()
// 	// 	panic(err)
// 	// }
// 	//
// 	// password, err := utils.HashPassword(cfg.Superadmin.Password)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	//
// 	// user, err := tx.User.Create().
// 	// 	SetFullName(cfg.Superadmin.Name).
// 	// 	SetEmail(cfg.Superadmin.Email).
// 	// 	SetPasswordHash(*password).
// 	// 	SetVerifiedAt(time.Now()).
// 	// 	Save(context.Background())
// 	// if err != nil {
// 	// 	_ = tx.Rollback()
// 	// 	panic(err)
// 	// }
// 	//
// 	// _, err = tx.UserRole.Create().
// 	// 	SetRoleID(role.ID).
// 	// 	SetProductID(product.ID).
// 	// 	SetUserID(user.ID).
// 	// 	Save(context.Background())
// 	// if err != nil {
// 	// 	_ = tx.Rollback()
// 	// 	panic(err)
// 	// }
// 	//
// 	// if err := tx.Commit(); err != nil {
// 	// 	return fmt.Errorf("failed to commit transaction: %w", err)
// 	// }
// 	//
// 	// fmt.Println("Seeder Super Admin created successfully")
//
// 	return nil
// }
