#!/usr/bin/env python3

import os
import threading

def gagag():
    os.system("go run go.go")

threading1 = threading.Thread(target=gagag)
threading1.daemon = True 
threading1.start()

os.system("python3 py.py")
