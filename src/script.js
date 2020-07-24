// https://codepen.io/liabru/pen/soGfr

// https://github.com/liabru/matter-js/blob/master/src/render/Render.js
Matter.Render.bodies = function (render, bodies, context) {
  var c = context,
    engine = render.engine,
    options = render.options,
    showInternalEdges = options.showInternalEdges || !options.wireframes,
    body,
    part,
    i,
    k

  for (i = 0; i < bodies.length; i++) {
    body = bodies[i]

    if (!body.render.visible) continue

    // handle compound parts
    for (k = body.parts.length > 1 ? 1 : 0; k < body.parts.length; k++) {
      part = body.parts[k]

      if (!part.render.visible) continue

      if (options.showSleeping && body.isSleeping) {
        c.globalAlpha = 0.5 * part.render.opacity
      } else if (part.render.opacity !== 1) {
        c.globalAlpha = part.render.opacity
      }

      if (
        part.render.sprite &&
        part.render.sprite.texture &&
        !options.wireframes
      ) {
        // part sprite
        var sprite = part.render.sprite,
          texture = _getTexture(render, sprite.texture)

        c.translate(part.position.x, part.position.y)
        c.rotate(part.angle)

        c.drawImage(
          texture,
          texture.width * -sprite.xOffset * sprite.xScale,
          texture.height * -sprite.yOffset * sprite.yScale,
          texture.width * sprite.xScale,
          texture.height * sprite.yScale
        )

        // revert translation, hopefully faster than save / restore
        c.rotate(-part.angle)
        c.translate(-part.position.x, -part.position.y)
      } else {
        // part polygon
        if (part.circleRadius) {
          c.beginPath()
          c.arc(
            part.position.x,
            part.position.y,
            part.circleRadius,
            0,
            2 * Math.PI
          )
        } else {
          c.beginPath()
          c.moveTo(part.vertices[0].x, part.vertices[0].y)

          for (var j = 1; j < part.vertices.length; j++) {
            if (!part.vertices[j - 1].isInternal || showInternalEdges) {
              c.lineTo(part.vertices[j].x, part.vertices[j].y)
            } else {
              c.moveTo(part.vertices[j].x, part.vertices[j].y)
            }

            if (part.vertices[j].isInternal && !showInternalEdges) {
              c.moveTo(
                part.vertices[(j + 1) % part.vertices.length].x,
                part.vertices[(j + 1) % part.vertices.length].y
              )
            }
          }

          c.lineTo(part.vertices[0].x, part.vertices[0].y)
          c.closePath()
        }

        if (!options.wireframes) {
          c.fillStyle = part.render.fillStyle

          if (part.render.lineWidth) {
            c.lineWidth = part.render.lineWidth
            c.strokeStyle = part.render.strokeStyle
            c.stroke()
          }

          c.fill()
        } else {
          c.lineWidth = 1
          c.strokeStyle = '#bbb'
          c.stroke()
        }
      }

      c.globalAlpha = 1

      //Here's the custom part
      if (part.render.text) {
        //30px is default font size
        var fontsize = 30
        //arial is default font family
        var fontfamily = part.render.text.family || 'Arial'
        //white text color by default
        var color = part.render.text.color || '#FFFFFF'

        if (part.render.text.size) fontsize = part.render.text.size
        else if (part.circleRadius) fontsize = part.circleRadius / 2

        var content = ''
        if (typeof part.render.text == 'string') content = part.render.text
        else if (part.render.text.content) content = part.render.text.content

        c.textBaseline = 'middle'
        c.textAlign = 'center'
        c.fillStyle = color
        c.font = fontsize + 'px ' + fontfamily
        c.fillText(content, part.position.x, part.position.y)
      }
    }
  }
}

var Example = Example || {}
var w = window.innerWidth
var h = window.innerHeight
const content = document.querySelector('.taskball-physics')
const spritesArea = document.querySelector('.taskball-sprites')

class Taskball {
  constructor(id) {
    var x = Math.random() * w
    var y = Math.random() * -h
    var base = w / 30
    if (base < 5) base = 5
    if (base > 10) base = 10
    var multiplier = w / 10
    if (multiplier < 30) multiplier = 30
    if (multiplier > 100) multiplier = 100

    this.radius = base + Math.random() * multiplier
    this.body = Matter.Bodies.circle(x, y, this.radius, {
      // density: 0.0005,
      // frictionAir: 0.06,
      // restitution: 0.3,
      // friction: 0.01,
      render: {
        fillStyle: 'black',
        text: {
          content: 'プログラミング学習',
          color: 'white',
          size: 10,
          family: 'Papyrus',
        },
      },
      label: 'id:' + id,
    })

    this.element = document.createElement('div')
    this.element.className =
      'taskball ' + 'taskball--' + Math.floor(Math.random() * 5)
    this.element.style.width = this.radius * 2 + 'px'
    this.element.style.height = this.radius * 2 + 'px'
    spritesArea.appendChild(this.element)
  }

  update() {
    this.pos = { x: this.body.position.x, y: this.body.position.y }
    this.element.style.transform = `translate(${
      this.pos.x - this.radius - 8
    }px, ${this.pos.y - this.radius - 8}px)`
  }
}

const taskballs = []
const numCircle = 30
for (var i = 0; i < numCircle; i++) {
  taskballs.push(new Taskball(i))
}

Example.avalanche = function () {
  var Engine = Matter.Engine,
    Render = Matter.Render,
    Runner = Matter.Runner,
    Composite = Matter.Composite,
    Composites = Matter.Composites,
    Common = Matter.Common,
    MouseConstraint = Matter.MouseConstraint,
    Mouse = Matter.Mouse,
    World = Matter.World,
    Bodies = Matter.Bodies

  var engine = Engine.create(),
    world = engine.world

  var render = Render.create({
    element: content,
    engine: engine,
    options: {
      width: w,
      height: h,
      wireframes: false,
    },
  })

  Render.run(render)

  var runner = Runner.create()
  Runner.run(runner, engine)

  var balls = taskballs.map((taskball) => taskball.body)
  World.add(world, balls)

  ground = Matter.Bodies.rectangle(w / 2, h + 30, w, 60, {
    isStatic: true,
  })
  wall1 = Matter.Bodies.rectangle(-30, h / 2, 60, h * 2, { isStatic: true })
  wall2 = Matter.Bodies.rectangle(w + 30, h / 2, 60, h * 2, {
    isStatic: true,
  })
  World.add(world, [ground, wall1, wall2])

  var mouse = Mouse.create(render.canvas)
  var mouseConstraint = MouseConstraint.create(engine, {
    mouse: mouse,
    constraint: {
      stiffness: 0.2,
      render: {
        visible: true,
      },
    },
  })
  World.add(world, mouseConstraint)
  render.mouse = mouse

  Matter.Events.on(mouseConstraint, 'mousemove', (event) => {
    var foundPhysics = Matter.Query.point(balls, event.mouse.position)
    console.log(foundPhysics[0])
  })

  Matter.Events.on(engine, 'afterUpdate', () => {
    if (mouseConstraint.body) {
      // console.info(mouseConstraint.body)
    }
  })

  // fit the render viewport to the scene
  //   Render.lookAt(render, Composite.allBodies(world))

  return {
    engine: engine,
    runner: runner,
    render: render,
    canvas: render.canvas,
    stop: function () {
      Matter.Render.stop(render)
      Matter.Runner.stop(runner)
    },
  }
}

Example.avalanche()
