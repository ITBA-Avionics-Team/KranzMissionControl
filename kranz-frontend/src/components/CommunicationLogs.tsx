import React, { useContext, useEffect, useState } from 'react';
import { SystemStatusContext } from '../contexts/SystemStatusContext';
import { SystemStatus } from '../model/SystemStatus';

{

}

const CommunicationLogs = () => {
  const { latestSystemStatus, setLatestSystemStatus } = useContext<SystemStatus>(SystemStatusContext);

  return (
    <div>
      {latestSystemStatus?.launchpad?.current_state}
    </div>
  );
};

export default CommunicationLogs;
