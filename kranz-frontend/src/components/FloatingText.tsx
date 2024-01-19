import React, { useState } from 'react';

const FloatingText = ({cursorPos, text}) => {

  return (
      <div
        style={{
          position: 'absolute',
          left: cursorPos.x,
          top: cursorPos.y,
          pointerEvents: 'none', // Prevent the text from interfering with other mouse events
        }}
      >
        {text}
      </div>
  );
};

export default FloatingText;
