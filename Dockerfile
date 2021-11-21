FROM golang
RUN apt update && apt install vim tmux curl wget git -y