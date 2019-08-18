# Keyboard Backlight control for macbooks.

Keyboard-backlight settings for lubuntu on my macbook pro, 13 inch retina early 2015

## Use

`keyboard-backlight --up`
`keyboard-backlight --down`

## Notes

This works on my macbook, it may cause yours to say mean things and slam shut on your fingers.

I also had to

`chown root:root keyboard-backlight`
`chmod u+s keyboard-backlight`

in order for the command to have to rights to change the backlight setting.

In Lubuntu, I've assigned keyboard shortcuts in Configuration Center|Shortcut Keys:

* XF86kbdBrightnessDown set to `keyboard-backlight --down`
* XF86kbdBrightnessUp set to `keyboard-backlight --up`
