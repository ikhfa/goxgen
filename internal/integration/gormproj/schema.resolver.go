package gormproj

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/goxgen/goxgen/internal/integration/gormproj/generated"
	"github.com/goxgen/goxgen/utils"
	"gorm.io/gorm/clause"
)

// NewUser is the resolver for the new_user field.
func (r *mutationResolver) NewUser(ctx context.Context, input *generated.NewUser) (*generated.User, error) {
	cars := make([]*generated.Car, len(input.Cars))
	for i, car := range input.Cars {
		cars[i] = &generated.Car{
			ID:   utils.Deref(car.ID),
			Make: utils.Deref(car.Make),
			Done: utils.Deref(car.Done),
		}
	}
	user := &generated.User{
		Name: input.Name,
		Cars: cars,
	}
	res := r.DB.Preload(clause.Associations).Create(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

// DeleteUsers is the resolver for the delete_users field.
func (r *mutationResolver) DeleteUsers(ctx context.Context, input *generated.DeleteUsers) ([]*generated.User, error) {
	var users []*generated.User
	r.DB.Delete(&users, input.Ids)
	return users, nil
}

// NewCar is the resolver for the new_car field.
func (r *mutationResolver) NewCar(ctx context.Context, input *generated.CarInput) (*generated.Car, error) {
	panic(fmt.Errorf("not implemented: NewCar - new_car"))
}

// UpdateCar is the resolver for the update_car field.
func (r *mutationResolver) UpdateCar(ctx context.Context, input *generated.CarInput) (*generated.Car, error) {
	panic(fmt.Errorf("not implemented: UpdateCar - update_car"))
}

// XgenIntrospection is the resolver for the _xgen_introspection field.
func (r *queryResolver) XgenIntrospection(ctx context.Context) (*generated.XgenIntrospection, error) {
	return r.Resolver.XgenIntrospection()
}

// ListCars is the resolver for the list_cars field.
func (r *queryResolver) ListCars(ctx context.Context, input *generated.ListCars) ([]*generated.Car, error) {
	var cars []*generated.Car
	res := r.DB.Where(&[]*generated.Car{
		{
			ID:     utils.Deref(input.ID),
			UserID: utils.Deref(input.UserID),
		},
	}).Find(&cars)

	if res.Error != nil {
		return nil, res.Error
	}
	return cars, nil
}

// ListUser is the resolver for the list_user field.
func (r *queryResolver) ListUser(ctx context.Context, input *generated.ListUser) ([]*generated.User, error) {
	var users []*generated.User
	res := r.DB.Where(&[]*generated.User{
		{
			ID:   utils.Deref(input.ID),
			Name: utils.Deref(input.Name),
		},
	}).Find(&users)

	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
