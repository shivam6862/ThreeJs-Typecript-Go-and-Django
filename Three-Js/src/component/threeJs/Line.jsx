import React from "react";
import * as three from "three";
import { OrbitControls } from "three/examples/jsm/controls/OrbitControls";
import { gsap } from "gsap";
import { useEffect } from "react";
import { useRef } from "react";

const Line = () => {
  const canvasRef = useRef(null);
  useEffect(() => {
    const canvas = canvasRef.current;
    const renderer = new three.WebGLRenderer({ canvas });
    renderer.setSize(window.innerWidth, window.innerHeight);
    const camera = new three.PerspectiveCamera(
      45,
      window.innerWidth / window.innerHeight,
      1,
      100
    );
    camera.position.set(0, 0, 100);
    camera.lookAt(0, 0, 0);

    const scene = new three.Scene();
    const material = new three.LineBasicMaterial({ color: 0xff66ff });
    const points = [];
    points.push(new three.Vector3(0, 0, 0));
    points.push(new three.Vector3(-12, 0, 0));
    points.push(new three.Vector3(0, 12, 0));
    points.push(new three.Vector3(12, 0, 0));
    points.push(new three.Vector3(0, -12, 0));
    points.push(new three.Vector3(-12, 0, 0));

    const geometry = new three.BufferGeometry().setFromPoints(points);
    const line = new three.Line(geometry, material);
    scene.add(line);
    renderer.render(scene, camera);
  }, []);

  return (
    <>
      <canvas ref={canvasRef}></canvas>
    </>
  );
};
export default Line;
