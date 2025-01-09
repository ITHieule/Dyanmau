package repo_impl

type UserReponImpl struct{
	spl*db.Sql
}

 func NewUserRepo(sql *db.Sql) reponsitory.UserRepo{
	return &UserReponImpl {
		sql:sql,
	}
 }
 func ( u UserReponImpl) SaveUser(context context.Context,user model.User)(model.User , error){
	statement := '
	INSERT INTO users(user_id,emai,password,role,full_name,created_at,updated_at)
	VALUE (:user_id, :email, :password,: role, :full_name, :created_at, :updated_at)
'
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := u.sql.Db.NamedExeContext(context,statement,user)
	if err != nil{
		log.Error(err.Error())
		if err, ok:= err.(*pq/Error);ok{
			if err.Code.Name()=="unique_violation"{
				return user,banana.UserConflick
			}
		}
		return user,banana.SignUpFail
	}
	return user,nil
 }