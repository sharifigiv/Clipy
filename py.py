from pydub import AudioSegment
from pydub.playback import play

sound = AudioSegment.from_file('sounds/Clip.wav')

f = open('data/clip.txt', 'r')
old = len(f.readlines())
f.close()

while True:
    f = open('data/clip.txt', 'r')

    new = len(f.readlines())

    if new != old:
        old = new
        play(sound)

    f.close()


        