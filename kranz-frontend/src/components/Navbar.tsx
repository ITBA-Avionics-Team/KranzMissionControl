import { useContext } from "react";
import { SendCommandModalContext } from "../contexts/SendCommandModalContext";

const Navbar = () => {
  const { showSendCommandModal, setShowSendCommandModal } = useContext<boolean>(SendCommandModalContext);

  const navbarStyle = {
    width: '100vw',
    height: '5rem',
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',

    backgroundColor: 'var(--main-bg-color)',
    borderColor: 'var(--main-fg-color)',
    color: 'var(--main-fg-color)',
    borderBottomStyle: 'solid',
    borderWidth: '1px'
  }

  return (
    <div style={navbarStyle}>
      <h3 style={{marginLeft: '2rem'}}>Kranz Mission Control 🚀</h3>
      <button
        onClick={() => setShowSendCommandModal(true)}>
        Send Command
      </button>
    </div>
  );
};

export default Navbar;