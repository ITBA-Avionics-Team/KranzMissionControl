const Navbar = () => {

  const navbarStyle = {
    width: '100vw',
    height: '5rem',
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',

    backgroundColor: 'var(--main-bg-color)',
    borderColor: 'var(--main-fg-color)',
    color: 'var(--main-fg-color)',
    borderBottomStyle: 'solid',
    borderWidth: '1px'
  }

  return (
    <div style={navbarStyle}>
      <h3>Kranz Mission Control 🚀</h3>
    </div>
  );
};

export default Navbar;