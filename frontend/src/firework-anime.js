let Animation = {}

Animation.setup = function () {
  Animation.canvas = document.querySelector('#fireworks')

  var ctx = Animation.canvas.getContext('2d')

  Animation.setCanvasSize = function () {
    Animation.canvas.width = window.innerWidth * 2
    Animation.canvas.height = window.innerHeight * 2
    Animation.canvas.style.width = window.innerWidth + 'px'
    Animation.canvas.style.height = window.innerHeight + 'px'
    Animation.canvas.getContext('2d').scale(2, 2)
  }

  Animation.setParticuleDirection = function (p) {
    var angle = (anime.random(0, 360) * Math.PI) / 180
    var value = anime.random(50, 180)
    var radius = [-1, 1][anime.random(0, 1)] * value
    return {
      x: p.x + radius * Math.cos(angle),
      y: p.y + radius * Math.sin(angle),
    }
  }

  Animation.createParticule = function (x, y) {
    var colors = ['#FF1461', '#18FF92', '#5A87FF', '#FBF38C']
    var p = {}
    p.x = x
    p.y = y
    p.color = colors[anime.random(0, colors.length - 1)]
    p.radius = anime.random(16, 32)
    p.endPos = Animation.setParticuleDirection(p)
    p.draw = function () {
      ctx.beginPath()
      ctx.arc(p.x, p.y, p.radius, 0, 2 * Math.PI, true)
      ctx.fillStyle = p.color
      ctx.fill()
    }
    return p
  }

  Animation.createCircle = function (x, y) {
    var p = {}
    p.x = x
    p.y = y
    p.color = '#FFF'
    p.radius = 0.1
    p.alpha = 0.5
    p.lineWidth = 6
    p.draw = function () {
      ctx.globalAlpha = p.alpha
      ctx.beginPath()
      ctx.arc(p.x, p.y, p.radius, 0, 2 * Math.PI, true)
      ctx.lineWidth = p.lineWidth
      ctx.strokeStyle = p.color
      ctx.stroke()
      ctx.globalAlpha = 1
    }
    return p
  }

  Animation.renderParticule = function (anim) {
    for (var i = 0; i < anim.animatables.length; i++) {
      anim.animatables[i].target.draw()
    }
  }

  Animation.animateParticules = function (x, y) {
    var circle = Animation.createCircle(x, y)
    var particules = []
    var numberOfParticules = 30

    for (var i = 0; i < numberOfParticules; i++) {
      particules.push(Animation.createParticule(x, y))
    }
    anime
      .timeline()
      .add({
        targets: particules,
        x: function (p) {
          return p.endPos.x
        },
        y: function (p) {
          return p.endPos.y
        },
        radius: 0.1,
        duration: anime.random(600, 800),
        easing: 'easeOutExpo',
        update: Animation.renderParticule,
        complete: function (anim) {
          console.log('hoge')
          Animation.hideCanvas()
        },
      })
      .add(
        {
          targets: circle,
          radius: anime.random(80, 160),
          lineWidth: 0,
          alpha: {
            value: 0,
            easing: 'linear',
            duration: anime.random(600, 800),
          },
          duration: anime.random(1200, 1800),
          easing: 'easeOutExpo',
          update: Animation.renderParticule,
          // offset: 0
        },
        0
      )
  }

  Animation.hideCanvas = function () {
    Animation.canvas.setAttribute('style', 'display: none')
  }

  Animation.showCanvas = function () {
    Animation.canvas.setAttribute('style', '')
  }

  const render = anime({
    duration: Infinity,
    update: function () {
      ctx.clearRect(0, 0, Animation.canvas.width, Animation.canvas.height)
    },
  })

  Animation.explode = function (x, y) {
    render.play()
    Animation.showCanvas()
    Animation.animateParticules(x, y)
  }

  // window.human = true

  // Animation.updateCoords = function (e) {
  //   x = e.clientX || e.touches[0].clientX
  //   y = e.clientY || e.touches[0].clientY
  //   return {x: x, y: y}
  // }

  // var tap =
  //   'ontouchstart' in window || navigator.msMaxTouchPoints
  //     ? 'touchstart'
  //     : 'mousedown'

  // document.addEventListener(
  //   tap,
  //   function (e) {
  //     window.human = true
  //     render.play()
  //     const xy = Animation.updateCoords(e)
  //     Animation.animateParticules(xy.x, xy.y)
  //   },
  //   false
  // )

  // function autoClick() {
  //   var centerX = window.innerWidth / 2
  //   var centerY = window.innerHeight / 2

  //   if (window.human) return
  //   animateParticules(
  //     anime.random(centerX - 50, centerX + 50),
  //     anime.random(centerY - 50, centerY + 50)
  //   )
  //   anime({ duration: 200 }).finished.then(autoClick)
  // }

  // autoClick()
  // setCanvasSize();
  // window.addEventListener('resize', setCanvasSize, false);

  Animation.hideCanvas()
}

Animation.setup()
