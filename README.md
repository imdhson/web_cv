# web_cv
[![Demo on YouTube](https://img.youtube.com/vi/wLmLqxQizKw/0.jpg)](https://www.youtube.com/watch?v=wLmLqxQizKw)

<br>

[한국어 README.md](README-ko.md)
<br>
A simple way to using OpenCV function on the web.<br>
using python-OpenCV, Golang WebServer.<br>
additionaly, simple implementation of ajax json transfer.
<br>
If server is ON, try here: http://webcv.imdhs.one
<br> but server status is NOT guaranteed. I don't have money.
## Quick execution:
### Ubuntu, Debian

    sudo apt install -y golang;  sudo apt install -y python; pip install numpy; pip install opencv-python; pip install matplotlib; pip install cvlib;git clone https://github.com/imdhson/web_cv.git ; cd web_cv/webserver/ ; go build; ./webcv
 
### Fedora, Centos, RedHat
    
    sudo yum install -y golang;  sudo yum install -y python; pip install numpy; pip install opencv-python; pip install matplotlib; pip install cvlib;git clone https://github.com/imdhson/web_cv.git ; cd web_cv/webserver/ ; go build; ./webcv
    
### macOS
    
    brew install golang;  brew install python; pip install numpy; pip install opencv-python; pip install matplotlib; pip install cvlib;git clone https://github.com/imdhson/web_cv.git ; cd web_cv/webserver/ ; go build; ./webcv
    
### Windows
#### Install GoLang, python manually. and type below.

    pip install numpy; pip install opencv-python; pip install matplotlib; pip install cvlib;git clone https://github.com/imdhson/web_cv.git ; cd web_cv/webserver/ ; go build; ./webcv
    
----

## Start sequencly:

### 1. Install the GoLang, python.

#### Ubuntu, Debian

    sudo apt install -y golang;  sudo apt install -y python;
  
#### Fedora, CentOS, RedHat

    sudo yum install -y golang;  sudo yum install -y python
  
#### macOS

    sudo brew install golang;  sudo brew install python

#### Windows
##### [GoLang install](https://go.dev/dl/)
##### [Python install](https://www.python.org/downloads/)
    
### 2. Install the dependency packages.

    pip install numpy; pip install opencv-python; pip install matplotlib; pip install cvlib
  
### 3. Clone and build
  
    git clone https://github.com/imdhson/web_cv.git ; cd web_cv/webserver/ ; go build; ./webcv
  
### 4. Start with your modern webbrowser.

    http://localhost:[port you selected]

Daegu University Computer Enginnering &copy; imdhson,&copy;  Leehyunjoon

Opensoure usage: OpenCV, Golang, html5up.net
