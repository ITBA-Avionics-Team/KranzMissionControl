import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import CommunicationLogs from './components/CommunicationLogs'
import { SystemStatusProvider } from './providers/SystemStatusProvider'
import SystemVisualizer from './components/SystemVisualizer'
import FloatingText from './components/FloatingText'
import Navbar from './components/Navbar'

function App() {
  
  return (
    <div style={{ height: '100vh', width: '100vw' , display:'flex', flexDirection: 'column'}}>
      <SystemStatusProvider>
        <Navbar />
        <SystemVisualizer />
        {/* <CommunicationLogs /> */}
      </SystemStatusProvider>
    </div>
  )
}

export default App
