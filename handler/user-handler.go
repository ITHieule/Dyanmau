package handler

import (
	"duanmau/database"
	"github.com/labstack/echo/v4"
	"net/http"
	validator"github.com/go-playground/validator/v10"
security""
	uuid"github.com/google/uuid"
)
type UserHandler struct{
	UserRepo reponsitory.UserRepo

}

func (u* Dangnhap) Dangnhap(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "lan",
		"email": "lan@gmail.com",
	})
}
func(u* Dangki)  Dangki(c echo.Context) error {
req := req2.ReqDangki{}
	if err = c.Bind(&req); err != nil{
		log.Error(err.Error())
		return c.JSON(http.StatusBabRequest,model.Reponse{
			StatusCode: http.StatusBabRequest,
			Message: err.Error,
			Data: nil,
		})
	}
validate := validate.New()

	if err:= vilidator.Struct(req)err != nil{
		log.Error(err.Error())
		return c.JSON(http.StatusBabRequest,model.Reponse{
			StatusCode: http.StatusBabRequest,
			Message: err.Error,
			Data: nil,
		})
	}

	hash := security.HashAndSalt([]byte (req.Password))
	role := model.MEMBER.String() 

	userId, err := uuid.NewUUID()
	if err != nil{
		log.Error(err.Error())
		return c.JSON(http.StatusBabRequest,model.Reponse{
			StatusCode: http.StatusBabRequest,
			Message: err.Error,
			Data: nil,
		})
	}

	user := model.User{
		UserId : userId,
	FullName  req.FullName,
	Email req.Email
	Password  hash,
	Role rloe,
	Token "",
}	
	user,err = u.UserRepo.SaveUser(c.Request().Context(),)
if err != nil{
	return c.JSON(http.StatusConflict,model.Reponse{
		StatusCode: http.StatusConflict,
			Message: err.Error,
			Data: nil,
	})
}
	return c.JSON(http.StatusOK, model.Reponse{
		StatusCode: http.StatusConflict,
			Message: err.Error,
			Data: user,
	})
}
func (u* GetUsers) GetUsers(c echo.Context) error {
	// Truy vấn dữ liệu từ database
	query := "SELECT id, title, completed FROM todos"
	rows, err := database.DB.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":   "Không thể truy vấn dữ liệu",
			"details": err.Error(),
		})
	}
	defer rows.Close()

	// Duyệt qua kết quả và xây dựng danh sách todos
	var todos []map[string]interface{}
	for rows.Next() {
		var (
			id        int
			title     string
			completed bool
		)
		if err := rows.Scan(&id, &title, &completed); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error":   "Lỗi khi đọc dữ liệu",
				"details": err.Error(),
			})
		}
		todos = append(todos, echo.Map{
			"id":        id,
			"title":     title,
			"completed": completed,
		})
	}

	// Kiểm tra nếu không có dữ liệu
	if len(todos) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Không tìm thấy công việc nào",
		})
	}

	// Trả về danh sách todos dưới dạng JSON
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Danh sách công việc",
		"data":    todos,
	})
}