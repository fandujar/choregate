import { Admin } from "@/components/Admin"
import { Header } from "@/components/Header"
import { Menu } from "@/components/Menu"

export const AdminPage = () => {
    return (
        <div>
            <Header/>
            <div className="flex">
                <Menu />
                <Admin />
            </div>
        </div>
    )
}