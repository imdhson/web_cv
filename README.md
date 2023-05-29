# web_cv
A simple way to using OpenCV function on the web. <br>
<br>
##How to execute the server:
###on general unix system:
 git clone https://github.com/imdhson/web_cv.git ; cd web_cv/webserver;./webcv
###if not:
1. Install the GoLang, python.
#### Ubuntu, Debian
  sudo apt install -y golang;  sudo apt install -y python;
#### Fedora, Centos
  sudo yum install -y golang;  sudo yum install -y python
2. Install the dependency packages.
  pip install numpy; pip install opencv-python; pip install matplotlib; pip install cvlib
4. clone and build
  git clone https://github.com/imdhson/web_cv.git ; cd web_cv ; go build; ./webcv

&copy; 대구대학교 컴퓨터공학과 손동휘, 이현준

