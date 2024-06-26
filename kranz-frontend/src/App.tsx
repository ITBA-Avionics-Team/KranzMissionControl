import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import SystemStatusText from './components/SystemStatusText'
import { SystemStatusProvider } from './providers/SystemStatusProvider'
import SystemVisualizer from './components/SystemVisualizer'
import FloatingText from './components/FloatingText'
import Navbar from './components/Navbar'
import CustomCommandSender from './components/CustomCommandSender/CustomCommandSender'
import SendCommandModal from './components/SendCommandModal/SendCommandModal'
import { SendCommandModalProvider } from './providers/SendCommandModalProvider'

function App() {
  
  return (
    <div style={{ height: '100vh', width: '100vw' , display:'flex', flexDirection: 'column'}}>
      <SystemStatusProvider>
        <SendCommandModalProvider>
          <Navbar />
          <SystemVisualizer />
          <SendCommandModal />
        </SendCommandModalProvider>
      </SystemStatusProvider>
    </div>
  )
}

export default App
