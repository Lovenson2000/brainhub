
import Sidebar from "@/components/Sidebar";
import React from "react";

export default function Page() {
  return (
    <div className="flex">
      <Sidebar />
      <h1 className="text-2xl font-bold">Welcome to the Dashboard</h1>
    </div>
  );
}
