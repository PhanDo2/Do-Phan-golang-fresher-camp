Khoá ngoại không được khuyến khích dùng là bởi vì một số yếu điểm sau :
- 1 hệ thống lớn thì được chia làm các service nhỏ mà mỗi service nhỏ thì thường có 1 DB riêng, nếu có FK tồn tại thì liên kết sẽ bị gẫy hầu hết sẽ đi qua Res API của nhau 
- FK sẽ ảnh hưởng đến tốc độ INSERT, UPDATE, DELETE bởi vì FK yêu cầu dàng buộc dữ liệu 1 cách nhất quán 
- Dữ liệu phải tải lại toàn bồ khi thêm 1 bảng mới từ bên ngoài vào hoặc phải theo 1 cái khung cũ khó khăn cho việc update
