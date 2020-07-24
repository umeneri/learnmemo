// https://codepen.io/liabru/pen/soGfr

var Example = Example || {}
var w = window.innerWidth
var h = window.innerHeight
const content = document.querySelector('.eyeball-physics')
const spritesArea = document.querySelector('.eyeball-sprites')

class Eyeball {
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
      render: {
        fillStyle: '#C44D58',
        text: {
          content: 'Test',
          color: 'blue',
          size: 16,
          family: 'Papyrus',
        },
      },
      label: 'id:' + id,
    })

    this.element = document.createElement('div')
    this.element.className =
      'eyeball ' + 'eyeball--' + Math.floor(Math.random() * 5)
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

  // add bodies
  // var stack = Composites.stack(0, 0, 10, 5, 0, 0, function(x, y) {
  //     return Bodies.circle(x, y, Common.random(10, 50), { friction: 0.00001, restitution: 0.5, density: 0.001 });
  // });

  // World.add(world, stack);

  const eyeballs = []
  const numCircle = 3
  for (var i = 0; i < numCircle; i++) {
    eyeballs.push(new Eyeball(i))
  }
  World.add(
    world,
    eyeballs.map((eyeball) => eyeball.body)
  )
  ground = Matter.Bodies.rectangle(w / 2, h + 30, w, 60, {
    isStatic: true,
  })
  wall1 = Matter.Bodies.rectangle(-30, h / 2, 60, h * 2, { isStatic: true })
  wall2 = Matter.Bodies.rectangle(w + 30, h / 2, 60, h * 2, {
    isStatic: true,
  })
  World.add(world, [ground, wall1, wall2])

  var mouse = Mouse.create(render.canvas)
  // mouse.pixelRatio = pixelDensity()
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

  Matter.Events.on(engine, 'afterUpdate', () => {
    if (mouseConstraint.body) {
      console.info(mouseConstraint.body)
    }

    // eyeballs.forEach((eye) => {
    //   eye.update()
    // })
  })

  // fit the render viewport to the scene
  //   Render.lookAt(render, Composite.allBodies(world))

  // wrapping using matter-wrap plugin
  //   for (var i = 0; i < stack.bodies.length; i += 1) {
  //     stack.bodies[i].plugin.wrap = {
  //       min: { x: render.bounds.min.x, y: render.bounds.min.y },
  //       max: { x: render.bounds.max.x, y: render.bounds.max.y },
  //     }
  //   }

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

function render() {
  var bodies = Matter.Composite.allBodies(this.engine.world)
  var canvas = document.getElementById('canvas')
  var context = canvas.getContext('2d')

  window.requestAnimationFrame(this.render)

  context.fillStyle = '#FFFFFF'
  context.fillRect(0, 0, canvas.width, canvas.height)
  context.globalAlpha = 1
  context.beginPath()

  for (var i = 0; i < bodies.length; i += 1) {
    var part = bodies[i]

    if (part.render.text) {
      var fontsize = 30
      var fontfamily = part.render.text.family || 'Arial'
      var color = part.render.text.color || '#FF0000'

      if (part.render.text.size) {
        fontsize = part.render.text.size
      } else if (part.circleRadius) {
        fontsize = part.circleRadius / 2
      }

      var content = ''
      if (typeof part.render.text === 'string') {
        content = part.render.text
      } else if (part.render.text.content) {
        content = part.render.text.content
      }

      context.fillStyle = 'black'
      context.save()
      context.translate(part.position.x, part.position.y)

      context.textBaseline = 'middle'
      context.textAlign = 'center'
      context.fillStyle = color
      context.font = fontsize + 'px ' + fontfamily
      context.fillText(content, 0, 0)
      context.restore()
      context.fillStyle = 'blue'
      context.fillRect(part.position.x, part.position.y, 10, 10)
    }
    var vertices = bodies[i].vertices
    context.moveTo(vertices[0].x, vertices[0].y)

    for (var j = 1; j < vertices.length; j += 1) {
      context.lineTo(vertices[j].x, vertices[j].y)
    }

    context.lineTo(vertices[0].x, vertices[0].y)
  }

  context.lineWidth = 1.5
  context.strokeStyle = '#000000'
  context.stroke()
}


Example.avalanche()
