import React, { createContext, useState, useContext } from 'react';

export const SendCommandModalContext = createContext<any>({showModal: true, currentCommand:{}});
