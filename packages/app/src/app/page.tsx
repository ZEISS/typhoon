import { Metadata } from 'next'
import DefaultLayout from '@/components/default-layout'
import { PropsWithChildren } from 'react'
import { newClient } from '@/api'

export const metadata: Metadata = {
  title: 'Typhoon',
  description: ''
}

async function getTeams() {
  return await newClient().team.listTeam()
}

export interface RootProps extends PropsWithChildren { }

export default async function Root({ children }: Readonly<RootProps>) {
  return (
    <>
      <DefaultLayout>{children}</DefaultLayout>
    </>
  )
}
