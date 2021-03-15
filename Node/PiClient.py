def rainbow_hearts():
    import requests
    url = "http://3.140.244.63:8081/check"  # "http://192.168.0.190:8081/check"

    from sense_hat import SenseHat
    import time

    sense = SenseHat()
    sense.low_light = True

    # Create variables to hold each RGB color we want to use

    # red
    r = (255, 0, 0)

    # pink
    p = (204, 0, 204)

    # orange
    o = (255, 128, 0)

    # yellow
    y = (255, 255, 0)

    # green
    g = (0, 255, 0)

    # aqua
    a = (0, 255, 255)

    # blue
    b = (0, 0, 255)

    # purple
    pr = (128, 0, 255)

    # empty (no color)
    e = (0, 0, 0)

    red_heart = [
        e, e, e, e, e, e, e, e,
        e, r, r, e, r, r, e, e,
        r, r, r, r, r, r, r, e,
        r, r, r, r, r, r, r, e,
        r, r, r, r, r, r, r, e,
        e, r, r, r, r, r, e, e,
        e, e, r, r, r, e, e, e,
        e, e, e, r, e, e, e, e
    ]

    pink_heart = [
        e, e, e, e, e, e, e, e,
        e, p, p, e, p, p, e, e,
        p, p, p, p, p, p, p, e,
        p, p, p, p, p, p, p, e,
        p, p, p, p, p, p, p, e,
        e, p, p, p, p, p, e, e,
        e, e, p, p, p, e, e, e,
        e, e, e, p, e, e, e, e
    ]

    orange_heart = [
        e, e, e, e, e, e, e, e,
        e, o, o, e, o, o, e, e,
        o, o, o, o, o, o, o, e,
        o, o, o, o, o, o, o, e,
        o, o, o, o, o, o, o, e,
        e, o, o, o, o, o, e, e,
        e, e, o, o, o, e, e, e,
        e, e, e, o, e, e, e, e
    ]
    heart_colors = [red_heart, pink_heart, orange_heart]
    i = True
    sense.clear()
    while i:
        for color in heart_colors:
            try:
                if requests.get(url).content == "1":
                    sense.set_pixels(color)
                else:
                    sense.clear()
                time.sleep(1)
            except:
                print "Could not connect"


rainbow_hearts()
