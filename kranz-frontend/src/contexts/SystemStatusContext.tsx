import React, { createContext, useState, useContext } from 'react';
import { DefaultSystemStatus, SystemStatus } from '../model/SystemStatus';

export const SystemStatusContext = createContext<SystemStatus>(DefaultSystemStatus);