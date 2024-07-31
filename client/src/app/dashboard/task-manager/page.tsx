import React, { ComponentType } from "react";
import dynamic from "next/dynamic";

const UserDetails: ComponentType<any> = dynamic(
  () =>
    import("../../../components/UserDetails.tsx").then((mod) => mod.default),
  {
    ssr: false,
  }
);

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
