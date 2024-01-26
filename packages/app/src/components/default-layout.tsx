import '@/styles/globals.css'
import { type ReactNode } from 'react'
import { ThemeToggle } from '@/components/theme-toggle'
import { UserNav } from '@/components/user-nav'
import { SystemSwitcher } from '@/components/system-switcher'
import { TeamSwitcher } from './team-switcher'
import { SlashIcon } from '@radix-ui/react-icons'
import { MainNav } from '@/components/main-nav'

interface DefaultLayoutProps {
  children: ReactNode
}

export default async function DefaultLayout({
  children
}: Readonly<DefaultLayoutProps>) {
  return (
    <>
      <div className="flex-row">
        <div className="border-b">
          <div className="flex h-16 items-center px-4">
            <div className="flex items-center flex-1 space-x-4">
              <TeamSwitcher />
              <SlashIcon width={30} height={30} />
              <SystemSwitcher />

              <MainNav />
              <ThemeToggle />
              <UserNav />
            </div>
          </div>
        </div>
        <div className="flex-1">{children}</div>
      </div>
    </>
  )
}
