import { BrowserRouter } from 'react-router-dom';

import { Router } from './Router';
import { AuthProvider } from './hooks/Auth';

function App() {
  return (
    <div className='bg-slate-200 text-slate-950 h-screen'>
    <BrowserRouter>
    <AuthProvider>
      <Router />
    </AuthProvider>
    </BrowserRouter>
    </div>
  )
}

export default App
