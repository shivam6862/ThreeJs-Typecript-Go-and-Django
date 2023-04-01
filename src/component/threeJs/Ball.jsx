import React from "react";
import * as three from "three";
import { OrbitControls } from "three/examples/jsm/controls/OrbitControls";
import { gsap } from "gsap";
import { useEffect } from "react";
import { useRef } from "react";

const Ball = () => {
  const canvasRef = useRef(null);
  useEffect(() => {
    // scene
    const scene = new three.Scene();
    const geometry = new three.SphereGeometry(3, 64, 64);
    const material = new three.MeshStandardMaterial({
      color: "#00ff85",
      roughness: 0.5,
    });
    const mesh = new three.Mesh(geometry, material);
    scene.add(mesh);

    //   sizes
    const sizes = {
      width: window.innerWidth,
      height: window.innerHeight,
    };

    //   light
    const light = new three.PointLight(0xffffff, 1, 100);
    light.position.set(0, 10, 10);
    light.intensity = 1;
    scene.add(light);

    //   camara
    const camera = new three.PerspectiveCamera(
      45,
      sizes.width / sizes.height,
      0.1,
      100
    );
    camera.position.z = 20;
    scene.add(camera);

    //   render
    const canvas = canvasRef.current;
    const renderer = new three.WebGLRenderer({ canvas });
    renderer.setSize(sizes.width, sizes.height);
    renderer.setPixelRatio(2);
    renderer.render(scene, camera);

    //   controls
    const controls = new OrbitControls(camera, canvas);
    controls.enableDamping = true;
    controls.enablePan = false;
    controls.enableZoom = false;
    controls.autoRotate = true;
    controls.autoRotateSpeed = 5;

    //   resize
    window.addEventListener("resize", () => {
      // update sizes
      sizes.width = window.innerWidth;
      sizes.height = window.innerHeight;
      // update camara
      camera.aspect = sizes.width / sizes.height;
      camera.updateProjectionMatrix();
      renderer.setSize(sizes.width, sizes.height);
    });

    //   loop and axis rotation
    var a = 1;
    const loop = () => {
      mesh.position.x += 0.01 * a;
      if (mesh.position.x >= (sizes.width / sizes.height) * 5) a = -1;
      if (mesh.position.x < 0) a = 1;
      controls.update();
      renderer.render(scene, camera);
      window.requestAnimationFrame(loop);
    };
    loop();

    //   animation for scaling
    const tl = gsap.timeline({ defaults: { duration: 1 } });
    tl.fromTo(mesh.scale, { z: 0, x: 0, y: 0 }, { z: 1, y: 1, x: 1 });

    //   mouse animation color
    let mouseDown = false;
    let rgb = [];
    window.addEventListener("mousedown", () => (mouseDown = true));
    window.addEventListener("mouseup", () => (mouseDown = false));
    window.addEventListener("mousemove", (e) => {
      if (mouseDown) {
        rgb = [
          Math.round((e.pageX / sizes.width) * 255),
          Math.round((e.pageX / sizes.height) * 255),
          150,
        ];
        //   lets animate
        let newColor = new three.Color(`rgb(${rgb.join(",")})`);
        gsap.to(mesh.material.color, {
          r: newColor.r,
          g: newColor.g,
          b: newColor.b,
        });
      }
    });
  }, []);

  return (
    <>
      <canvas ref={canvasRef}></canvas>
    </>
  );
};
export default Ball;
