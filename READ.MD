# STEP 1 ติดตั้ง gRPC go version > 1.31
``` $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest ```
 
# ตรวจสอบว่า protoc และ protoc-gen-go-grpc ถูกติดตั้งหรือไม่
``` $ protoc --version ```
``` $ protoc-gen-go-grpc --version ```

# ถ้าคำสั่งเหล่านี้ไม่ทำงาน อาจต้องติดตั้ง gRPC Plugin ด้วย:
``` $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest ```

# STEP 2 Map Env GO (MAC)
``` $ export PATH="$PATH:$(go env GOPATH)/bin" ```

# STEP 3
``` $ go mod tidy ```

# ล้าง cache mod
``` go clean -modcache ```

# backend-challenge ทั้งหมด 3 ข้อ อยู่ที่ Folder ชื่อ Module สามารถดูได้จากไฟล์ main.go