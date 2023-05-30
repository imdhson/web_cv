import cv2
import sys
path = sys.argv[1]
img = cv2.imread (path, 0)
cv2.imwrite(path, img)