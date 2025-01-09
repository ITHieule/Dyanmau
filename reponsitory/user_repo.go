package reponsitory

type User interface{
	SaveUser(context context.Context,user model.User)(model.User , error)
}