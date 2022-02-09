# Go ex 10
Vì sao không nên chứa file upload vào ngay chính bên trong service mà nên dùng Cloud.
  - Bởi vì khi upload ảnh vào ngay chính bên trong serice thì data base phải lưu trữ các file vật lý dẫn đến sẽ rất nặng nề về dung lượng => ảnh hưởng rất lớn đến hiệu năng cũng như là tốc độ load dữ liệu của serviec 
  - Do DB có giới hạn chính vì vậy lưu vào Cloud thì sẽ lưu được nhiều hơn, tiện hơn trong việc lấy dữ liệu.
  
Vì sao không chứa binary ảnh vào DB?
  - Để tăng tính bảo mật
