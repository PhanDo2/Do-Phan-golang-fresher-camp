# Do-Phan-golang-fresher-camp
Khi nào cần tạo các cột số đếm ngay trên table dữ liệu (VD: liked_count trên restaurants)?
  - Khi mà API chịu tải quá cao(VD: tải read trên API list của restaurant, food) thì chúng ta cần tạo các cột số đếm ngay trên table dữ liệu. Bởi vì khi đó dữ liệu của các bảng rất lớn dẫn đến việc đi xuống DB để đếm sẽ kém hiệu quả do cần Db phải tính toán lại.
