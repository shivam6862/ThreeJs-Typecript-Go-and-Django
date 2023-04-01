import React from "react";
import * as three from "three";
import { OrbitControls } from "three/examples/jsm/controls/OrbitControls";
import { gsap } from "gsap";
import { useEffect } from "react";
import { useRef } from "react";

const Text = () => {
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
    //
    renderer.render(scene, camera);
  }, []);

  return (
    <>
      <canvas ref={canvasRef}></canvas>
    </>
  );
};
export default Text;
