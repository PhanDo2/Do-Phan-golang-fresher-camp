Tác dụng của khoá chính PK
- Để dữ liệu trên các dòng trùng nhau 
- Nếu 1 bảng trong DB có nhiều hơn 1 thuộc tính là PK thì Pk nào đứng trước thì được ưu tiến sắp xếp trước(sắp xếp dựa trên vị trí vật lý của Pk đứng đầu)
- thường được dùng để truy xuất dữ liệu một cách nhanh nhất 
Vì sao dùng ID là số tự tăng? 
- Để khoá chính luôn là duy nhất thì chúng ta cần sử dụng công thức để thêm khoá chính. Một cách đơn giản trong giải pháp đó là ta để khoá chính tự động tăng như vậy giảm tiểu tối đa được lỗi ràng buộc của PK
khoá chính Pk dùng trong nhiều cột 
- Khi chúng ta cần nối dữ liệu giữa các bảng (thường là 2 bảng) 
- Khi bảng đó không mang ý nghĩa là 1 thực thể 
