"use client";
import { ComponentType } from "react";
import dynamic from "next/dynamic";

const UserDetails: ComponentType<any> = dynamic(
  () =>
    import("../../components/UserDetails").then((mod) => mod.default),
  {
    ssr: false,
  }
);

export default function Page() {
  return (
    <div className="p-4">
      <UserDetails />
    </div>
  );
}
