# Do-Phan-golang-fresher-camp
Trong trường hợp tạo cột đếm thì làm sao để update cột đó? Làm sao để API chính không bị block vì phải update số đếm?
  - Để update được cột đếm thì ta sẽ tạo 2 Api: user_like_restaurant và user_unlike_restaurant trong module restaurant_like
  - Để tránh khi đang update bị block API chính thì trường hợp này ta không thể dùng cách update cứng được mà nên để cho Db tự update 
      (VD :db.Model(&product).UpdateColumn("like_count", gorm.Expr("like_count - ?", 1))// UPDATE "restaurantlike" SET "like_count" = like_count - 1 WHERE "id" = 3;) và nên để logic này chạy ở goroutine riêng 
  
