import * as readline from 'node:readline/promises';
import * as fs from 'node:fs';
import { stdin, stdout } from 'node:process';

const rl = readline.createInterface(stdin, stdout);
stdin.isTTY && stdin.setRawMode(true);

async function yesNo(prompt) {
    const result = await rl.question(`* 是否添加${prompt}：y=是，n=否？ `) === 'y';
    stdout.moveCursor(0, -1);
    stdout.clearLine();
    return result;
}

async function makeList(value, values) {
    return await rl.question(`-> ${value}: [${values.join(', ')}] （默认：${values[0]}）： `) || values[0];
}

async function makeString(value, defaultValue) {
    const result = defaultValue ?
        await rl.question(`-> ${value} （默认：${defaultValue}）： `) :
        await rl.question(`-> ${value}: `);
    return result || defaultValue;
}

async function makeNumber(value, defaultValue) {
    const result = defaultValue ?
        await rl.question(`-> ${value} （默认：${defaultValue}）： `) :
        await rl.question(`-> ${value}: `);
    return Number(result || defaultValue) || undefined;
}

async function makeBool(value, defaultValue) {
    const result = defaultValue ?
        await rl.question(`-> ${value} y=是，n=否 （默认：${defaultValue}）： `) :
        await rl.question(`-> ${value} y=是，n=否： `);
    return (result || defaultValue) === 'y' || undefined;
}

async function run() {
    while (true) {
        const item = await createItem();
        if (!item) { continue; }
        const resultJson = JSON.stringify(item, null, 2);
        stdout.write(`\r\n************* ${item.name} *************\r\n\r\n${resultJson}\r\n\r\n`);
        item.stats ? storeNewUnit(item) : storeNewItem(item);
    }
}

function storeNewUnit(unit) {
    const configPath = 'units_config.json';
    let configStr = fs.readFileSync(configPath);
    const config = JSON.parse(configStr);
    config.push(unit);
    configStr = JSON.stringify(config, null, 4);
    fs.writeFileSync(configPath, configStr);
}

function storeNewItem(item) {
    const configPath = 'items_config.json';
    let configStr = fs.readFileSync(configPath);
    const config = JSON.parse(configStr);
    switch (item.type) {
        case 'armor': config.armor.push(item); break;
        case 'weapon': config.weapon.push(item); break;
        case 'magic': config.magic.push(item); break;
        case 'disposable': config.disposable.push(item); break;
        case 'ammunition': config.ammunition.push(item); break;
        case 'provision': config.provision.push(item); break;
        default: return;
    }
    configStr = JSON.stringify(config, null, 4);
    fs.writeFileSync(configPath, configStr);
}

async function createItem() {
    const r = await rl.question(
        `选择类型：
        1 - 护甲（armor）
        2 - 武器（weapon）
        3 - 魔法（magic）
        4 - 消耗品（disposable）
        5 - 弹药（ammunition）
        6 - 食物（provision）
        7 - 单位（unit）\r\n`);
    const i = Number.parseInt(r);
    switch (i) {
        case 1:
            return await makeArmor();
        case 2:
            return await makeWeapon();
        case 3:
            return await makeMagic();
        case 4:
            return await makeDisposable();
        case 5:
            return await makeAmmunition();
        case 6:
            return await makeProvision();
        case 7:
            return await makeUnit();
        default:
            return null;
    }
}

async function makeUnit() {
    stdout.write('单位：\r\n');
    const result = {
        name: await makeString('名称（name）'),
        code: await makeString('代码（code）'),
        code: await makeString('职业（class）'),
        stats: await makeUnitStats('属性面板（stats）'),
        slots: await makeSlots('装备栏位（slots）'),
        booty: await makeUnitBooty('战利品（booty）'),
        achievements: await makeAchievements('成就（achievements）'),
        inventory: await makeInventory('背包（inventory）'),
        drop: await makeDrop('掉落（drop）'),
    }
    return result;
}

async function makeArmor() {
    stdout.write('护甲：\r\n');
    const result = {
        ...await makeItem('armor'),
        ...await makeEquipment(),
    }
    return result;
}

async function makeWeapon() {
    stdout.write('武器：\r\n');
    const result = {
        ...await makeItem('weapon'),
        ammunitionKind: await makeString('弹药种类（ammunitionKind）'),
        range: await makeActionRange('作用范围（range）'),
        useCost: await makeUnitBaseAttributes('使用消耗（useCost）'),
        ...await makeEquipment(),
        damage: [],
    }
    while (await yesNo('伤害（damage）')) {
        result.damage.push(await makeDamageImpact('伤害（damage）'));
    }
    return result;
}

async function makeMagic() {
    stdout.write('魔法：\r\n');
    const result = {
        ...await makeItem('magic'),
        requirements: await makeUnitRequirements('使用要求（requirements）'),
        range: await makeActionRange('作用范围（range）'),
        useCost: await makeUnitBaseAttributes('使用消耗（useCost）'),
        damage: [],
        modification: [],
    }
    while (await yesNo('伤害（damage）')) {
        result.damage.push(await makeDamageImpact('伤害（damage）'));
    }
    while (await yesNo('属性效果（modification）')) {
        result.modification.push(await makeUnitModificationImpact('属性效果（modification）'));
    }
    return result;
}

async function makeDisposable() {
    stdout.write('消耗品：\r\n');
    const result = {
        ...await makeItem('disposable'),
        range: await makeActionRange('作用范围（range）'),
        useCost: await makeUnitBaseAttributes('使用消耗（useCost）'),
        damage: [],
        modification: [],
    }
    while (await yesNo('伤害（damage）')) {
        result.damage.push(await makeDamageImpact('伤害（damage）'));
    }
    while (await yesNo('属性效果（modification）')) {
        result.modification.push(await makeUnitModificationImpact('属性效果（modification）'));
    }
    return result;
}

async function makeAmmunition() {
    stdout.write('弹药：\r\n');
    const result = {
        ...await makeItem('ammunition'),
        kind: await makeString('种类（kind）'),
        range: await makeActionRange('作用范围（range）'),
        damage: [],
    }
    while (await yesNo('伤害（damage）')) {
        result.damage.push(await makeDamageImpact('伤害（damage）'));
    }
    return result;
}

async function makeProvision() {
    stdout.write('食物：\r\n');
    const result = {
        ...await makeItem('provision'),
        recovery: await makeUnitRecovery('恢复（recovery）')
    }
    return result;
}

async function makeUnitStats(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        progress: await makeUnitProgress('成长（progress）'),
        baseAttributes: await makeUnitBaseAttributes('基础属性（baseAttributes）'),
        attributes: await makeUnitAttributes('属性（attributes）'),
        resistance: await makeDamage('抗性（resistance）'),
    }
    return result;
}

async function makeUnitProgress(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        level: await makeNumber('等级（level）', 1),
        experience: await makeNumber('经验（experience）'),
    }
    return result;
}

async function makeInventory(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        descriptor: [],
    }
    while (await yesNo('物品（item）')) {
        result.descriptor.push(await makeItemDescriptor('物品（item）'));
    }
    return result;
}

async function makeItemDescriptor(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        code: await makeString('代码（code）'),
        quantity: await makeNumber('数量（quantity）'),
        equipped: await makeBool('是否已装备（equipped）'),
    }
    return result;
}

async function makeSlots(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        head: await makeNumber('头部（head）'),
        neck: await makeNumber('颈部（neck）'),
        body: await makeNumber('身体（body）'),
        hand: await makeNumber('手部（hand）'),
        leg: await makeNumber('腿部（leg）'),
        weapon: await makeNumber('武器（weapon）'),
    }
    return result;
}

async function makeAchievements(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {};
    while (await yesNo('成就（achievement）')) {
        const key = await makeString('成就标识（achievement id）');
        const value = await makeNumber('获得次数（times acquired）', 1);
        result[key] = value;
    }
    return result;
}

async function makeActionRange(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        minimumX: await makeNumber('最小横向距离（minimumX）'),
        maximumX: await makeNumber('最大横向距离（maximumX）'),
        minimumY: await makeNumber('最小纵向距离（minimumY）'),
        maximumY: await makeNumber('最大纵向距离（maximumY）'),
        radius: await makeNumber('半径（radius）'),
    }
    return result;
}

async function makeUnitBooty(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        coins: await makeNumber('金币（coins）'),
        ruby: await makeNumber('红宝石（ruby）'),
    }
    return result;
}

async function makeDrop(header) {
    // header && stdout.write(`${header}:\r\n`);
    const result = [];
    while (await yesNo('掉落（drop）')) {
        result.push({
            ...await makeUnitBooty('掉落（drop）'),
            weight: await makeNumber('权重（weight）', 1),
        });
    }
    return result;
}

async function makeItem(type, header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        type: type,
        name: await makeString('名称（name）'),
        description: await makeString('描述（description）'),
        code: await makeString('代码（code）'),
        canBeThrownAway: await makeBool('是否可丢弃（canBeThrownAway）'),
        canBeSold: await makeBool('是否可出售（canBeSold）'),
        price: await makeUnitBooty('价格（price）'),
    }
    return result;
}

async function makeEquipment(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        durability: await makeNumber('耐久上限（durability）', 500),
        wearout: await makeNumber('磨损（wearout）'),
        slot: await makeList('装备栏位（slot）', ['weapon', 'head', 'neck', 'body', 'hand', 'leg']),
        slotsNumber: await makeNumber('占用栏位数量（slotsNumber）', 1),
        requirements: await makeUnitRequirements('使用要求（requirements）'),
        modification: [],
    }
    while (await yesNo('属性效果（modification）')) {
        result.modification.push(await makeUnitModification('属性效果（modification）'));
    }
    return result;
}

async function makeUnitAttributes(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        strength: await makeNumber('力量（strength）'),
        physique: await makeNumber('体质（physique）'),
        agility: await makeNumber('敏捷（agility）'),
        endurance: await makeNumber('耐力（endurance）'),
        intelligence: await makeNumber('智力（intelligence）'),
        initiative: await makeNumber('先攻（initiative）'),
        luck: await makeNumber('幸运（luck）'),
    }
    return result;
}

async function makeUnitRequirements(header) {
    const result = {
        ...await makeUnitAttributes(header),
        class: await makeString('职业（class）'),
    }
    return result;
}

async function makeUnitBaseAttributes(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        health: await makeNumber('生命值（health）'),
        stamina: await makeNumber('体力（stamina）'),
        mana: await makeNumber('法力（mana）'),
        actionPoints: await makeNumber('行动点（actionPoints）', 9),
    }
    return result;
}

async function makeImpact(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        duration: await makeNumber('持续时间（duration）'),
        chance: await makeNumber('概率（chance）'),
        deviation: await makeNumber('浮动值（deviation）'),
    }
    return result;
}

async function makeDamageImpact(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        ...await makeImpact(),
        ...await makeDamage(),
    }
    return result;
}

async function makeUnitModificationImpact(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        ...await makeImpact(),
        ...await makeUnitModification(),
    }
    return result;
}

async function makeDamage(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        stabbing: await makeNumber('穿刺（stabbing）'),
        cutting: await makeNumber('斩击（cutting）'),
        crushing: await makeNumber('钝击（crushing）'),
        fire: await makeNumber('火焰（fire）'),
        cold: await makeNumber('寒冷（cold）'),
        lightning: await makeNumber('闪电（lightning）'),
        acid: await makeNumber('酸蚀（acid）'),
        poison: await makeNumber('毒素（poison）'),
        exhaustion: await makeNumber('体力消耗（exhaustion）'),
        manaDrain: await makeNumber('法力汲取（manaDrain）'),
        bleeding: await makeNumber('流血（bleeding）'),
        fear: await makeNumber('恐惧（fear）'),
        curse: await makeNumber('诅咒（curse）'),
        madness: await makeNumber('疯狂（madness）'),
    }
    return result;
}

async function makeUnitState(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        health: await makeNumber('生命值（health）'),
        stamina: await makeNumber('体力（stamina）'),
        mana: await makeNumber('法力（mana）'),
        stress: await makeNumber('压力（stress）'),
    }
    return result;
}

async function makeUnitRecovery(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {
        ...await makeUnitState(),
        ...await makeDamage(),
    }
    return result;
}

async function makeUnitModification(header) {
    header && stdout.write(`${header}:\r\n`);
    const result = {};
    if (await yesNo('基础属性（baseAttributes）')) {
        result.baseAttributes = await makeUnitBaseAttributes('基础属性（baseAttributes）');
    }
    if (await yesNo('属性（attributes）')) {
        result.attributes = await makeUnitAttributes('属性（attributes）');
    }
    if (await yesNo('抗性（resistance）')) {
        result.resistance = await makeDamage('抗性（resistance）');
    }
    if (await yesNo('伤害（damage）')) {
        result.damage = await makeDamage('伤害（damage）');
    }
    if (await yesNo('恢复（recovery）')) {
        result.recovery = await makeUnitRecovery('恢复（recovery）');
    }
    return result;
}

run();
