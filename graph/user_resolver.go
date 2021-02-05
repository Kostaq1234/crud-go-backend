package graph

type userResolver struct{ *Resolver }

//func (u *userResolver) Meetups(ctx context.Context,obj *models.User)([]*models.Meetup,error)  {
//	var m []*models.Meetup
//	for _,meetup:=range meetups{
//
//	}
//}
func (r *Resolver) User() *userResolver {

	return &userResolver{r}
}
