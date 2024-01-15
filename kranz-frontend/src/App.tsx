import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import CommunicationLogs from './components/CommunicationLogs'
import { SystemStatusProvider } from './providers/SystemStatusProvider'
import SystemVisualizer from './components/SystemVisualizer'

function App() {
  return (
    <>
      <SystemStatusProvider>
        <h1>Kranz Mission Control 🚀</h1>
        <p className="read-the-docs">
          Click on the Vite and React logos to learn more
        </p>
        <SystemVisualizer />
        <CommunicationLogs />
      </SystemStatusProvider>
    </>
  )
}

export default App
