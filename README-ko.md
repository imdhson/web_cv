# web_cv | ko-kr(한국어)
OpenCV 기능을 웹에서 쉽게 사용할 수 있도록 할게요.<br>
python-OpenCV, Go언어 웹서버를 이용합니다.<br>
추가로, json 전송을 ajax를 통해서 간단하게 구현해보았습니다.
<br>
만약 서버가 켜져있다면, 여기서 해보세요: http://webcv.imdhs.one <br>
그러나 서버 상태는 보장못합니다. 저는 거지에요. ㅠ.ㅠ
## 빠른 실행:
### Ubuntu, Debian

    sudo apt install -y golang;  sudo apt install -y python; pip install numpy; pip install opencv-python; pip install matplotlib; pip install cvlib;git clone https://github.com/imdhson/web_cv.git ; cd web_cv ; go build; ./webcv
 
### Fedora, Centos, RedHat
    
    sudo yum install -y golang;  sudo yum install -y python; pip install numpy; pip install opencv-python; pip install matplotlib; pip install cvlib;git clone https://github.com/imdhson/web_cv.git ; cd web_cv ; go build; ./webcv
    
----

## 한단계씩 설치:

### 1. 설치하기: GoLang, python.

#### Ubuntu, Debian

    sudo apt install -y golang;  sudo apt install -y python;
  
#### Fedora, Centos, RedHat

    sudo yum install -y golang;  sudo yum install -y python
  
### 2. 의존 패키지 설치하기

    pip install numpy; pip install opencv-python; pip install matplotlib; pip install cvlib
  
### 3. Clone과 빌드

    git clone https://github.com/imdhson/web_cv.git ; cd web_cv ; go build; ./webcv
  
### 4. 당신의 모던 웹브라우저로 시작하세요:

    http://localhost:[서버시작시 고른 포트]

대구대학교 컴퓨터공학과&copy;  손동휘,&copy;  이현준
