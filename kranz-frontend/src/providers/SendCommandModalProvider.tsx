import { useEffect, useState } from "react";
import { SendCommandModalContext } from "../contexts/SendCommandModalContext";
import { SystemStatusContext } from "../contexts/SystemStatusContext";
import { DefaultSystemStatus } from "../model/SystemStatus";

export const SendCommandModalProvider = ({ children }) => {
  const [showSendCommandModal, setShowSendCommandModal] = useState(false);

  return (
    <SendCommandModalContext.Provider value={{ showSendCommandModal, setShowSendCommandModal }}>
      {children}
    </SendCommandModalContext.Provider>
  );
};
