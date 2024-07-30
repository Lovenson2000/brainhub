
import React, { ComponentType } from "react";
import dynamic from "../../../../node_modules/next/dynamic";

const UserDetails: ComponentType = dynamic(() => import("../../../components/UserDetails"), {
  ssr: false,
});

export default function page() {
  return (
    <div>
      <h1>Task Manager</h1>
      <div>
        <UserDetails />
      </div>
    </div>
  );
}
