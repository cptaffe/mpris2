#!/usr/bin/env python3

from subprocess import call
import sys

usageString = """(c) 2015 Connor Taffe <cpaynetaffe@gmail.com>. All rights reserved.

Released under an MIT license (http://opensource.org/licenses/MIT).

usage:
mediaKeys [play | prev | next | vol-up | vol-down | mute]

play     signals play/pause keypress
prev     signals previous track keypress
next     signals next track keypress
vol-up   signals volume up keypress
vol-down signals voludme down keypress
mute     signals mute keypress"""

def usage():
	print(usageString)

def key(keyName):
	call(['xdotool', 'key', keyName]);

keyNames = {
	'play': 'XF86AudioPlay',
	'prev': 'XF86AudioPrev',
	'next': 'XF86AudioNext',
	# For completeness...
	'vol-up': 'XF86AudioRaiseVolume',
	'vol-down': 'XF86AudioLowerVolume',
	'mute': 'XF86AudioMute'
}

if len(sys.argv) == 2 and sys.argv[1] in keyNames:
	key(keyNames[sys.argv[1]])
else:
	usage()

