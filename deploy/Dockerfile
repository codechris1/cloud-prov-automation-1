# The default OS from this image is debian
FROM golang:latest

# Defining environment variables and Dockerfile parameters
ARG APP_NAME
ENV APP_NAME=${APP_NAME}

# Changing the default OS timezone, if you prefer UTC simply comment this line
# if you want other timezone pickup any of the timezones in folder /usr/share/zoneinfo/
# and replace it in the next line
RUN ln -sf /usr/share/zoneinfo/EST /etc/localtime

# Updating OS libraries and installing the required ones
RUN apt-get update -y
RUN apt-get install -y sudo less vim

# Copying the source code
WORKDIR /go
COPY cmd/${APP_NAME}/ src/${APP_NAME}/
COPY pkg/ src/
COPY configs/ configs/

# Compiling the packages
RUN go env -w GO111MODULE=off
RUN go install ${APP_NAME}

# Remove the source code from the image
RUN rm -Rf src/*

CMD ["sh","-c","${APP_NAME}"]