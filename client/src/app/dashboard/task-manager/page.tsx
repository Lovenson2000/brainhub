import React, { ComponentType } from "react";
import dynamic from "next/dynamic";
import TaskItem from "components/tasks/Task";

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
        <TaskItem />
      </div>
    </div>
  );
}
