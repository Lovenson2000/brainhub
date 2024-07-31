import Sidebar from "../../components/Sidebar";
import React from "react";

export interface Props {
  children?: React.ReactNode;
}

export default function DashboardLayout({ children }: Props) {
  return (
    <div className="dark:bg-slate-900 flex flex-col md:flex-row h-screen">
      <div className="w-full flex-none md:w-64">
        <Sidebar />
      </div>
      <div className="md:flex-grow p-4">{children}</div>
    </div>
  );
}
