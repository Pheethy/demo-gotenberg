# เริ่มต้นจากหยิบ golang image มาเป็น base image
FROM golang:1.23-alpine
# ทำการกำหนด /go/src/flavorparser เป็น path เริ่มต้น
WORKDIR /go/src/git.pheethy.co.th/demo-gotenberg
# Copy source ทั้งหมดจาก directory ปัจจุบันสู่ working directory ภายใน container
COPY . .
# ทำการ get complieDaemon สำหรับ run
RUN go get -u github.com/githubnemo/CompileDaemon
# ทำติดตั้ง complieDaemon สำหรับ run
RUN go install github.com/githubnemo/CompileDaemon
# download dependencies ทั้งหมดที่ใช้
RUN go mod tidy
RUN mkdir -p /go/src/git.pheethy.co.th/demo-gotenberg/tmp
# Build the Go app.
RUN go build -o /go/src/git.pheethy.co.th/demo-gotenberg/tmp/app main.go
# Expose port 3600 ออกมาภายนอก container
EXPOSE 3600
# กำหนดคำสั่งหลักที่จะรันเมื่อ container ถูกเรียกใช้งาน ในที่นี้คือเรียกใช้ compiledaemon
ENTRYPOINT CompileDaemon -include=go.mod -log-prefix=false -color=true -build="go build -o ./tmp/app main.go" -command="./tmp/app"
