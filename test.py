import random
import math

inside = 0
outside = 0

r = lambda: random.random()

for i in range(10000000000):
	x = r()
	y = r()
	u = (x*x)+(y*y)
	if u <= 1:
		inside +=1
	else:
		outside += 1
my_pi = 4*inside/(inside+outside)
err = abs(math.pi-my_pi)
print(my_pi, err)
print(math.pi/4)
