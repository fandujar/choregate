import { Menu } from "@/components/Menu"
import { Header } from "@/components/Header"
import { Dashboard } from "@/components/Dashboard"

export const DashboardPage = () => {
  return (
        <>
        <div>
          <Header/>
          <div className="flex">
            <Menu />
            <Dashboard />
          </div>
        </div>
        </>
    )
}