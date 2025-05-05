import React, { createContext, useState, useContext, Dispatch, SetStateAction } from 'react';
import { DefaultSystemStatus, SystemStatus } from '../model/SystemStatus';

interface SystemStatusContextType {
  latestSystemStatus: SystemStatus;
  setLatestSystemStatus: Dispatch<SetStateAction<SystemStatus>>;
}

const defaultContextValue: SystemStatusContextType = {
  latestSystemStatus: DefaultSystemStatus,
  setLatestSystemStatus: () => {},
};

export const SystemStatusContext = createContext<SystemStatusContextType>(defaultContextValue);