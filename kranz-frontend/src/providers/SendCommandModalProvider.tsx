import { useEffect, useState } from "react";
import { SendCommandModalContext } from "../contexts/SendCommandModalContext";
import { SystemStatusContext } from "../contexts/SystemStatusContext";
import { DefaultSystemStatus } from "../model/SystemStatus";

export const SendCommandModalProvider = ({ children }) => {
  const [sendCommandModalContext, setSendCommandModalContext] = useState({showModal: false, currentCommand:{}});

  return (
    <SendCommandModalContext.Provider value={{ sendCommandModalContext, setSendCommandModalContext }}>
      {children}
    </SendCommandModalContext.Provider>
  );
};
