import Link from "next/link";
import { GoTasklist } from "react-icons/go";
import { Button } from "./ui/button";
import Image from "next/image";

type SidebarItemProps = {
  label: string;
  icon: React.ReactNode;
  url: string;
};

export default function Sidebar() {
  return (
    <aside className="w-64 h-screen border-r">
      <div className="flex items-center gap-4 border border-r-0 px-4 py-2">
        <Image
          src="/assets/logo.svg"
          width={26}
          height={20}
          alt="logo"
          className="rounded-md shadow-md"
        />
        <h2 className="text-xl text-slate-600 font-medium">Brainhub</h2>
      </div>
      <div className="m-4 relative h-[90vh]">
        <nav className="mt-4">
          <SidebarItem
            label="Analytics"
            icon={<GoTasklist size={20} />}
            url="/analytics"
          />
          <SidebarItem
            label="Tasks"
            icon={<GoTasklist size={20} />}
            url="/tasks"
          />
          <SidebarItem
            label="Sessions"
            icon={<GoTasklist size={20} />}
            url="/sessions"
          />
          <SidebarItem
            label="Documents"
            icon={<GoTasklist size={20} />}
            url="/documents"
          />

          <SidebarItem
            label="Notes"
            icon={<GoTasklist size={20} />}
            url="/documents"
          />
        </nav>

        <Button
          type="submit"
          className="w-56 absolute bottom-4 bg-light-blue text-white py-2 rounded"
        >
          Go To Study Room
        </Button>
      </div>
    </aside>
  );
}

function SidebarItem({ label, icon, url }: SidebarItemProps) {
  return (
    <Link
      href={url}
      className="flex items-center py-2 text-slate-600 hover:text-light-blue hover:border hover:bg-white rounded-md transition-colors duration-200"
    >
      <span className="mr-2 font-semibold">{icon}</span>
      <span>{label}</span>
    </Link>
  );
}
