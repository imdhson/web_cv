import cv2
import sys

path = sys.argv[1]
img = cv2.imread (path)
gray_img = cv2.cvtColor (img, cv2.COLOR_BGR2GRAY)
cv2.imwrite(path, gray_img)