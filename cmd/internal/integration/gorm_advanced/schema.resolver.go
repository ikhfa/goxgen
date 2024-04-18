package gorm_advanced

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"github.com/ikhfa/goxgen/cmd/internal/integration/gorm_advanced/generated"
	"github.com/ikhfa/goxgen/plugins/cli/server"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

// UserBatchDelete is the resolver for the user_batch_delete field.
func (r *mutationResolver) UserBatchDelete(ctx context.Context, input *generated.DeleteUsers) ([]*generated.User, error) {
	var users []*generated.User
	r.DB.Delete(&users, input.Ids)
	return users, nil
}

// UserCreate is the resolver for the user_create field.
func (r *mutationResolver) UserCreate(ctx context.Context, input *generated.UserInput) (*generated.User, error) {
	u, err := input.ToUserModel(ctx)
	if err != nil {
		return nil, err
	}
	res := r.DB.Preload(clause.Associations).Create(u)
	if res.Error != nil {
		return nil, res.Error
	}
	return u, nil
}

// UserUpdate is the resolver for the user_update field.
func (r *mutationResolver) UserUpdate(ctx context.Context, input *generated.UserInput) (*generated.User, error) {
	u, err := input.ToUserModel(ctx)
	if err != nil {
		return nil, err
	}
	td := r.DB.Preload(clause.Associations).Save(u)
	if td.Error != nil {
		return nil, td.Error
	}
	return u, nil
}

// CarCreate is the resolver for the car_create field.
func (r *mutationResolver) CarCreate(ctx context.Context, input *generated.CarInput) (*generated.Car, error) {
	car, err := input.ToCarModel(ctx)
	if err != nil {
		return nil, err
	}
	res := r.DB.Preload(clause.Associations).Create(car)
	if res.Error != nil {
		return nil, res.Error
	}
	return car, nil
}

// CarUpdate is the resolver for the car_update field.
func (r *mutationResolver) CarUpdate(ctx context.Context, input *generated.CarInput) (*generated.Car, error) {
	car, err := input.ToCarModel(ctx)
	if err != nil {
		return nil, err
	}
	res := r.DB.Preload(clause.Associations).Save(car)
	if res.Error != nil {
		return nil, res.Error
	}
	return car, nil
}

// PhoneNumberCreate is the resolver for the phone_number_create field.
func (r *mutationResolver) PhoneNumberCreate(ctx context.Context, input *generated.PhoneNumberInput) (*generated.Phone, error) {
	p, err := input.ToPhoneModel(ctx)
	if err != nil {
		return nil, err
	}
	res := r.DB.Preload(clause.Associations).Create(p)
	if res.Error != nil {
		return nil, res.Error
	}
	return p, nil
}

// PhoneNumberUpdate is the resolver for the phone_number_update field.
func (r *mutationResolver) PhoneNumberUpdate(ctx context.Context, input *generated.PhoneNumberInput) (*generated.Phone, error) {
	p, err := input.ToPhoneModel(ctx)
	if err != nil {
		return nil, err
	}
	td := r.DB.Preload(clause.Associations).Save(p)
	if td.Error != nil {
		return nil, td.Error
	}
	return p, nil
}

// XgenIntrospection is the resolver for the _xgen_introspection field.
func (r *queryResolver) XgenIntrospection(ctx context.Context) (*generated.XgenIntrospection, error) {
	return generated.XgenIntrospectionValues()
}

// CarBrowse is the resolver for the car_browse field.
func (r *queryResolver) CarBrowse(ctx context.Context, where *generated.CarBrowseInput) ([]*generated.Car, error) {
	var cars []*generated.Car
	u, err := where.ToCarModel(ctx)
	if err != nil {
		return nil, err
	}
	res := r.DB.Preload(clause.Associations).Where(&[]*generated.Car{u}).Find(&cars)

	if res.Error != nil {
		return nil, res.Error
	}

	return cars, nil
}

// UserBrowse is the resolver for the user_browse field.
func (r *queryResolver) UserBrowse(ctx context.Context, where *generated.ListUser, pagination *generated.XgenPaginationInput, sort *generated.UserSortInput) ([]*generated.User, error) {
	// Get logger from context
	logger := server.GetLogger(ctx)
	logger.Info("UserBrowse", zap.Any("where", where))

	var users []*generated.User
	u, err := where.ToUserModel(ctx)
	if err != nil {
		return nil, err
	}
	res := r.DB.
		Preload(clause.Associations).
		Scopes(generated.Paginate(pagination), generated.Sort(sort)).
		Where(&[]*generated.User{u}).
		Find(&users)

	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

// PhoneNumberBrowse is the resolver for the phone_number_browse field.
func (r *queryResolver) PhoneNumberBrowse(ctx context.Context, where *generated.PhoneNumberBrowseInput) ([]*generated.Phone, error) {
	var phones []*generated.Phone
	u, err := where.ToPhoneModel(ctx)
	if err != nil {
		return nil, err
	}
	res := r.DB.Preload(clause.Associations).Where(&[]*generated.Phone{u}).Find(&phones)

	if res.Error != nil {
		return nil, res.Error
	}

	return phones, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
